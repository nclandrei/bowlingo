import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';


import { AppComponent } from './app.component';
import { ScoreComponent } from './components/score/score.component';
import { FormsModule } from '@angular/forms';
import { ScoreService } from './services/score.service';


@NgModule({
  declarations: [
    AppComponent,
    ScoreComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,
    NgbModule.forRoot()
  ],
  providers: [ScoreService],
  bootstrap: [AppComponent]
})
export class AppModule { }
