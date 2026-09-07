
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PredictionService } from '../../../services/Prediction.service';
import { Prediction } from '../../../models/Prediction';

@Component({
    selector: 'app-index-prediction',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPredictionComponent implements OnInit {

    predictions: Prediction[] = [];

    constructor(
        private router: Router,
        private service: PredictionService
) {}

    ngOnInit(): void {
        this.getPredictions();
}

    getPredictions(): void {
        this.service.getPredictions().subscribe((res) => {
        this.predictions = res;
    });
}

    deletePrediction(id: any): void {
        this.service.deletePrediction(id)
            .subscribe(() => {
                this.getPredictions();
            });
    }
}