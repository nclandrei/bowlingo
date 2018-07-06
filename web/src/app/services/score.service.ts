import { Injectable } from '@angular/core';
import { Frame } from '../models/frame.model';
import { Observable } from '../../../node_modules/rxjs/Observable';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class ScoreService {
  private scoreApiEndpoint = 'http://localhost:8000/api/score';

  constructor(private http: HttpClient) { }

  public sendFrames(frames: Frame[]): Observable<any> {
    return this.http.post(this.scoreApiEndpoint, {
      frames: frames
    });
  }
}
