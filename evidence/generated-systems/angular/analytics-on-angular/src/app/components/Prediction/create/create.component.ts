import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PredictionService } from '../../../services/Prediction.service';
import { Prediction } from '../../../models/Prediction';
import { SubBaseComponent } from '../../Prediction/sub.base.component';

@Component({
    selector: 'app-create-prediction',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePredictionComponent extends SubBaseComponent implements OnInit {

    title = 'Add Prediction';

    predictionForm: FormGroup;
    prediction: Prediction;

    constructor( http: HttpClient,
        private predictionService: PredictionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.predictionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  referenceKey: ['', Validators.required],
      predictedAt: ['', Validators.required],
      score: ['', Validators.required],
      Endpoint: ['', ],
      ModelVersion: ['', ],
      Dataset: ['', ]
        });
    }

    
    addPrediction(referenceKey, predictedAt, score, Endpoint, ModelVersion, Dataset): void {
        this.predictionService
        .addPrediction(referenceKey, predictedAt, score, Endpoint, ModelVersion, Dataset)
            .subscribe(() => {
                this.router.navigate(['/indexPrediction']);
            });
    }

    ngOnInit(): void {
    }
}