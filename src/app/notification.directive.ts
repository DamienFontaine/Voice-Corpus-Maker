import { Directive, ViewContainerRef } from '@angular/core';

@Directive({
  selector: '[appNotification]'
})
export class NotificationDirective {

  constructor(public viewContainerRef: ViewContainerRef) { }
}
