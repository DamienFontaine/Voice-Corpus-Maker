import { TestBed } from '@angular/core/testing';

import { HttpClientModule } from '@angular/common/http';

import { RecordService } from './record.service';

describe('RecordService', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientModule]
  }));

  it('should be created', () => {
    const service: RecordService = TestBed.get(RecordService);
    expect(service).toBeTruthy();
  });
});
