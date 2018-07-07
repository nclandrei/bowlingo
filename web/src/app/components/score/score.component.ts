import { Component, OnInit } from '@angular/core';
import { Frame } from '../../models/frame.model';
import { ScoreService } from '../../services/score.service';

@Component({
  selector: 'app-score',
  templateUrl: './score.component.html',
  styleUrls: ['./score.component.less']
})
export class ScoreComponent implements OnInit {

  private frames: Frame[] = [];
  private totalScore = 0;
  private firstRoll: number;
  private secondRoll: number;
  private bonusRoll: number;
  private error: any;

  constructor(private scoreService: ScoreService) { }

  ngOnInit() { }

  private sendScores() {
    const currentFrame: Frame = {
      firstRoll: this.firstRoll,
      secondRoll: this.secondRoll,
      bonusRoll: (this.frames.length === 9) ? this.bonusRoll : null,
    };
    this.frames.push(currentFrame);
    this.scoreService.sendFrames(this.frames)
      .subscribe((resp: any) => {
        this.totalScore = resp.score;
      },
      error => this.error = error);
  }

  private resetGame() {
    this.frames = [];
    this.totalScore = 0;
    this.error = null;
  }

  private getRollText(rollIndex: number, frame: Frame) {
    switch (rollIndex) {
      case 0:
        if (frame.firstRoll === 10) {
          return 'X';
        } else {
          return frame.firstRoll.toString();
        }
      case 1:
        if (frame.firstRoll === 10 && this.frames.length < 10) {
          return '\u00a0';
        } else if (frame.firstRoll + frame.secondRoll === 10) {
          return '/';
        } else if (frame.secondRoll === 10 && this.frames.length === 10) {
          return 'X';
        } else {
          return frame.secondRoll.toString();
        }
      case 2:
        if (this.frames.length < 10) {
          return '\u00a0';
        } else if (frame.bonusRoll === 10) {
          return 'X';
        } else {
          return frame.bonusRoll;
        }
    }
  }

  private isBonusRollAwarded(): boolean {
    return this.frames.length === 9 &&
      (this.frames[8].firstRoll === 10 || this.frames[8].firstRoll + this.frames[8].secondRoll === 10);
  }
}
