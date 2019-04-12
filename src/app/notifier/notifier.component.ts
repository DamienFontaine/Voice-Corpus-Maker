import { Component, OnInit, ViewChild, ComponentFactoryResolver, ComponentRef } from '@angular/core';
import { NotifierService } from '../notifier.service';

import { NotificationDirective } from '../notification.directive';
import { NotificationComponent } from '../notification/notification.component';

@Component({
  selector: 'app-notifier',
  templateUrl: './notifier.component.html',
  styleUrls: ['./notifier.component.less']
})
export class NotifierComponent implements OnInit {
  @ViewChild(NotificationDirective) notificationHost: NotificationDirective;

  components: ComponentRef<NotificationComponent>[] = [];

  constructor(private notifier: NotifierService, private componentFactoryResolver: ComponentFactoryResolver) {}

  ngOnInit() {
    this.getNotification();
  }

  getNotification(): void {
    this.notifier.get().subscribe(notification => {
      if (notification) {
        const viewContainerRef = this.notificationHost.viewContainerRef;
        const componentFactory = this.componentFactoryResolver.resolveComponentFactory(NotificationComponent);
        const componentRef = viewContainerRef.createComponent(componentFactory);
        (<NotificationComponent>componentRef.instance).notification = notification;
        (<NotificationComponent>componentRef.instance).ref = componentRef;
        this.components.push(componentRef);
        if (this.components.length > 4) {
          const c = this.components.shift();
          (<NotificationComponent>c.instance).close();
        }
      }
    });
  }
}
