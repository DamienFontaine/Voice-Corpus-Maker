import { Injectable } from '@angular/core';

import { Observable, Subject } from 'rxjs';
import { Notification } from './notification';

@Injectable({
  providedIn: 'root'
})
export class NotifierService {
  notification: Subject<Notification> = new Subject();

  constructor() { }

  notify(message, alert) {
    this.notification.next(new Notification(message, alert));
  }

  get(): Observable<Notification> {
    return this.notification;
  }
}
