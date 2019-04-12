import { Component, OnInit, ComponentRef } from '@angular/core';

import { Notification } from '../notification';
import { trigger, state, transition, animate, style } from '@angular/animations';
import { faWindowClose } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-notification',
  templateUrl: './notification.component.html',
  styleUrls: ['./notification.component.less'],
  animations: [
    trigger('openClose', [
      state('open', style({
        opacity: 1
      })),
      state('closed', style({
        opacity: 0
      })),
      transition('* => closed', [
        animate('0.5s')
      ]),
      transition('* => open', [
        animate('0.2s')
      ]),
    ]),
  ]
})
export class NotificationComponent implements OnInit {
  notification: Notification;
  ref: ComponentRef<NotificationComponent>;
  faWindowClose = faWindowClose;
  isOpen = true;

  constructor() { }

  ngOnInit() {}

  close() {
    this.isOpen = false;
    setTimeout(() => {
      this.ref.destroy();
    }, 1000);
  }
}
