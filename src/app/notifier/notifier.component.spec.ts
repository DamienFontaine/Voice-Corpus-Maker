import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NotifierComponent } from './notifier.component';
import { NotificationComponent } from '../notification/notification.component';

import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

describe('NotifierComponent', () => {
  let component: NotifierComponent;
  let fixture: ComponentFixture<NotifierComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports : [ FontAwesomeModule ],
      declarations: [ NotifierComponent, NotificationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NotifierComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
