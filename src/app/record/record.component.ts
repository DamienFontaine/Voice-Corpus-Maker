import { Component } from '@angular/core';

import { faMicrophone, faCircle, faPause } from '@fortawesome/free-solid-svg-icons';

import { NotifierService } from '../notifier.service';
import { Record } from '../record';
import { RecordService } from '../record.service';
import { Metadata } from '../metadata';
import { Recorder } from '../recorder';
import { Set } from '../set.enum';

declare const navigator: any;
declare const MediaRecorder: any;

@Component({
  selector: 'app-record',
  templateUrl: './record.component.html',
  styleUrls: ['./record.component.less']
})
export class RecordComponent  {
  mediaRecorder: any;
  recorder: Recorder;
  record: Record = new Record();

  faMicrophone = faMicrophone;
  faCircle = faCircle;
  faPause = faPause;

  isListening = false;
  chunks = [];
  context;

  constructor(private notifier: NotifierService, private recordService: RecordService) {
    this.record.metadata = new Metadata();
    this.record.metadata.set = Set.Train;
    if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
      navigator.mediaDevices.getUserMedia ({
        audio: true
      }).then(stream => {
        this.context = new AudioContext;
        this.mediaRecorder = new MediaRecorder(stream);
        const input = this.context.createMediaStreamSource(stream);
        this.recorder = new Recorder(input, { numChannels: 1 });
      }).catch(err => {
        console.log('The following getUserMedia error occured: ' + err);
      });
    } else {
      console.log('getUserMedia not supported on your browser!');
    }
  }

  listen() {
    if (this.isListening) {
      this.isListening = false;
      this.recorder.stop();
      this.recorder.exportWAV( (blob) => {
        this.record.file = new Blob([blob], { 'type' : 'audio/wav' });
        this.recordService.save(this.record).subscribe();
        this.recorder.clear();
      }, 'audio/wav');
    } else {
      if (!this.record.metadata.text) {
        this.notifier.notify('Text invalid', 'danger');
        return;
      }
      this.isListening = true;
      this.recorder.record();
    }
  }
}
