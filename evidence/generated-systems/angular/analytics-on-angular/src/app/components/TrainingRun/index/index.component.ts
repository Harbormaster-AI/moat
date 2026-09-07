
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TrainingRunService } from '../../../services/TrainingRun.service';
import { TrainingRun } from '../../../models/TrainingRun';

@Component({
    selector: 'app-index-trainingRun',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTrainingRunComponent implements OnInit {

    trainingRuns: TrainingRun[] = [];

    constructor(
        private router: Router,
        private service: TrainingRunService
) {}

    ngOnInit(): void {
        this.getTrainingRuns();
}

    getTrainingRuns(): void {
        this.service.getTrainingRuns().subscribe((res) => {
        this.trainingRuns = res;
    });
}

    deleteTrainingRun(id: any): void {
        this.service.deleteTrainingRun(id)
            .subscribe(() => {
                this.getTrainingRuns();
            });
    }
}