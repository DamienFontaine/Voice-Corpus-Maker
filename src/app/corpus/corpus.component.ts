import { Component, OnInit, ViewChild } from '@angular/core';

import { faHeadphones, faTrash, faFileDownload, faCircle } from '@fortawesome/free-solid-svg-icons';

import { RecordService } from '../record.service';
import { Record } from '../record';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
@Component({
  selector: 'app-corpus',
  templateUrl: './corpus.component.html',
  styleUrls: ['./corpus.component.less']
})
export class CorpusComponent implements OnInit {

  @ViewChild('video') player: any;
  source = '';

  current: number;
  limit = 5;
  pages: number[];

  faHeadphones = faHeadphones;
  faTrash = faTrash;
  faFileDownload = faFileDownload;
  faCircle = faCircle;

  records: Record[];

  constructor(private sanitizer: DomSanitizer, private recordService: RecordService) { }

  ngOnInit() {
    this.change(0);
  }

  getSource(): SafeUrl {
    return this.sanitizer.bypassSecurityTrustUrl(this.source);
  }

  change(page) {
    this.current = page;
    this.recordService.count().subscribe(count => {
      this.pages = Array(Math.ceil(count / this.limit)).fill('');
      if (this.current > this.pages.length - 1) {
        this.current = this.pages.length - 1;
      }
      this.recordService.findPerPage(this.current, this.limit).subscribe( records => {
        this.records = records;
      });
    });
  }

  delete(id) {
    this.recordService.delete(id).subscribe(res => {
      this.change(this.current);
    });
  }

  listen(id) {
    this.recordService.upload(id).subscribe(res => {
      this.source = URL.createObjectURL(res.body);
      this.player.nativeElement.load();
      this.player.nativeElement.play();
    });
  }

  export() {
    this.recordService.export().subscribe(res => {
      const blob = new Blob([res], { type: 'application/zip'});
      const url = window.URL.createObjectURL(blob);
      const downloadLink = document.createElement('a');
      downloadLink.href = url;
      document.body.appendChild(downloadLink);
      downloadLink.click();
      downloadLink.parentNode.removeChild(downloadLink);
    });
  }

  changeSet(id, set) {
    this.recordService.update(id, set).subscribe(res => {
      this.change(this.current);
    });
  }
}
