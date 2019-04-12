import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { NavbarComponent } from './navbar/navbar.component';
import { NotifierComponent } from './notifier/notifier.component';
import { NotificationDirective } from './notification.directive';
import { NotificationComponent } from './notification/notification.component';
import { RecordComponent } from './record/record.component';

import { CorpusComponent } from './corpus/corpus.component';
import { PagerComponent } from './pager/pager.component';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    NotifierComponent,
    NotificationDirective,
    NotificationComponent,
    RecordComponent,
    CorpusComponent,
    PagerComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FontAwesomeModule,
    HttpClientModule,
    FormsModule,
    BrowserAnimationsModule
  ],
  entryComponents: [
    NotificationComponent,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
