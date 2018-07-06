import { Component, OnInit } from '@angular/core';
import { Frame } from '../../models/frame.model';

@Component({
  selector: 'app-score',
  templateUrl: './score.component.html',
  styleUrls: ['./score.component.less']
})
export class ScoreComponent implements OnInit {

  private frames: Frame[] = [];
  private totalScore: number;
  private firstRoll: number;
  private secondRoll: number;

  constructor() {}

  ngOnInit() {}

  private sendScores() {
    console.log("called");
    const currentFrame: Frame = {
      FirstRoll: this.firstRoll,
      SecondRoll: this.secondRoll,
    };
    this.frames.push(currentFrame);
    console.log(this.frames);
  }
}
