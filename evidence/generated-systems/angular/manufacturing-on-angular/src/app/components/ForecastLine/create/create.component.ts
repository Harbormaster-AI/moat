import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ForecastLineService } from '../../../services/ForecastLine.service';
import { ForecastLine } from '../../../models/ForecastLine';
import { SubBaseComponent } from '../../ForecastLine/sub.base.component';

@Component({
    selector: 'app-create-forecastLine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateForecastLineComponent extends SubBaseComponent implements OnInit {

    title = 'Add ForecastLine';

    forecastLineForm: FormGroup;
    forecastLine: ForecastLine;

    constructor( http: HttpClient,
        private forecastLineService: ForecastLineService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.forecastLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  period: ['', Validators.required],
      quantity: ['', Validators.required],
      confidence: ['', Validators.required],
      Forecast: ['', ],
      Item: ['', ]
        });
    }

    
    addForecastLine(period, quantity, confidence, Forecast, Item): void {
        this.forecastLineService
        .addForecastLine(period, quantity, confidence, Forecast, Item)
            .subscribe(() => {
                this.router.navigate(['/indexForecastLine']);
            });
    }

    ngOnInit(): void {
    }
}