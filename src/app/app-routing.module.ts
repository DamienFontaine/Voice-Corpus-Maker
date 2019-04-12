import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { CorpusComponent } from './corpus/corpus.component';
import { RecordComponent } from './record/record.component';

const routes: Routes = [
  { path: 'corpus', component: CorpusComponent},
  { path: 'listen', component: RecordComponent },
  { path: '', redirectTo: 'listen', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
