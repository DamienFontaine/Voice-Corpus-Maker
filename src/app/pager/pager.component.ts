import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';
import { faAngleDoubleLeft, faAngleDoubleRight, faAngleLeft, faAngleRight } from '@fortawesome/free-solid-svg-icons';



@Component({
  selector: 'app-pager',
  templateUrl: './pager.component.html',
  styleUrls: ['./pager.component.less']
})
export class PagerComponent implements OnInit {
  @Output() changeRequest = new EventEmitter<number>();
  @Input() pages: number[];
  @Input() current: number;

  faAngleDoubleLeft = faAngleDoubleLeft;
  faAngleDoubleRight = faAngleDoubleRight;
  faAngleLeft = faAngleLeft;
  faAngleRight = faAngleRight;

  constructor() { }

  ngOnInit() {
  }

  change(page) {
    this.changeRequest.emit(page);
  }
}
