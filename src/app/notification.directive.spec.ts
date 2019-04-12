import { NotificationDirective } from './notification.directive';
import { Component, ViewChild, ViewContainerRef } from '@angular/core';
import { TestBed, async } from '@angular/core/testing';

@Component({
  template: `
  <div ref></div>
  `
})
export class MockComponent {
  @ViewChild('ref') ref: ViewContainerRef;
}

describe('NotificationDirective', () => {
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [MockComponent, NotificationDirective],
    }).compileComponents();
  }));

  it('should create an instance', () => {
    const fixture = TestBed.createComponent(MockComponent);
    const directive = new NotificationDirective(fixture.componentInstance.ref);
    expect(directive).toBeTruthy();
  });
});
