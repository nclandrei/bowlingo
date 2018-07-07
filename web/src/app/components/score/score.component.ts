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
  private bonusRoll = null;

  constructor(private scoreService: ScoreService) {}

  ngOnInit() {}

  private sendScores() {
    const currentFrame: Frame = {
      firstRoll: this.firstRoll,
      secondRoll: this.secondRoll,
      bonusRoll: this.bonusRoll,
    };
    this.frames.push(currentFrame);
    this.scoreService.sendFrames(this.frames)
      .subscribe((resp: any) => {
        this.totalScore = resp.score;
      });
  }

  private resetGame() {
    this.frames = [];
    this.totalScore = 0;
  }
}
