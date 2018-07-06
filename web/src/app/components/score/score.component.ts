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
  private totalScore: number;
  private firstRoll: number;
  private secondRoll: number;

  constructor(private scoreService: ScoreService) {}

  ngOnInit() {}

  private sendScores() {
    const currentFrame: Frame = {
      firstRoll: this.firstRoll,
      secondRoll: this.secondRoll,
      bonusRoll: null,
    };
    this.frames.push(currentFrame);
    this.scoreService.sendFrames(this.frames)
      .subscribe((resp: any) => {
        console.log(resp);
        this.totalScore = resp.score;
      });
  }

  private resetGame() {
    this.frames = [];
    this.totalScore = 0;
  }
}
