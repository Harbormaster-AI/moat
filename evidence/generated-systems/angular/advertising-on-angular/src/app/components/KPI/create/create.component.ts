import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { KPIService } from '../../../services/KPI.service';
import { KPI } from '../../../models/KPI';
import { SubBaseComponent } from '../../KPI/sub.base.component';

@Component({
    selector: 'app-create-kPI',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateKPIComponent extends SubBaseComponent implements OnInit {

    title = 'Add KPI';

    kPIForm: FormGroup;
    kPI: KPI;

    constructor( http: HttpClient,
        private kPIService: KPIService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.kPIForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  targetValue: ['', Validators.required],
      Campaign: ['', ],
      MetricType: ['', ]
        });
    }

    
    addKPI(targetValue, Campaign, MetricType): void {
        this.kPIService
        .addKPI(targetValue, Campaign, MetricType)
            .subscribe(() => {
                this.router.navigate(['/indexKPI']);
            });
    }

    ngOnInit(): void {
    }
}