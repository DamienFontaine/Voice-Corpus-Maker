import { Injectable } from '@angular/core';

import { HttpClient } from '@angular/common/http';

import { Record } from './record';

import { Observable, of } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class RecordService {

  constructor(private http: HttpClient) { }

  findAll(): Observable<Record[]> {
    return this.http.get<Record[]>('/record');
  }

  findPerPage(page: number, limit: number):  Observable<Record[]> {
    return this.http.get<Record[]>('/record/page/' + page + '/' + limit);
  }

  count(): Observable<number> {
    return this.http.get<number>('/record/count');
  }

  upload(id) {
    return this.http.get('/upload/' + id, {observe: 'response', responseType: 'blob' });
  }

  delete(id) {
    return this.http.delete('/record/' + id);
  }

  update(id , set) {
    const data = new FormData();
    data.append('set', set);
    return this.http.post('/record/' + id, data);
  }

  save(record: Record) {
    const data = new FormData();
    data.append('file', record.file);
    data.append('text', record.metadata.text);
    data.append('set', record.metadata.set);

    return this.http.post('/record', data);
  }

  export() {
    return this.http.get('/record/export', {responseType: 'arraybuffer'});
  }
}
