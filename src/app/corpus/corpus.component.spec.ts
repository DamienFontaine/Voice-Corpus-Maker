import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HttpClientModule } from '@angular/common/http';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { CorpusComponent } from './corpus.component';
import { PagerComponent } from '../pager/pager.component';


describe('CorpusComponent', () => {
  let component: CorpusComponent;
  let fixture: ComponentFixture<CorpusComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientModule, FontAwesomeModule],
      declarations: [ CorpusComponent, PagerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CorpusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
