import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InspectionPlanService } from '../../../services/InspectionPlan.service';
import { SubBaseComponent } from '../../InspectionPlan/sub.base.component';


@Component({
    selector: 'app-edit-inspectionPlan',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInspectionPlanComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InspectionPlan';

    inspectionPlanForm: FormGroup;
    inspectionPlan: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InspectionPlanService,
        private fb: FormBuilder
) {
        super(http);
        this.inspectionPlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  planNumber: ['', Validators.required],
      revision: ['', Validators.required],
      Item: ['', ],
      Characteristics: ['', ],
      SamplingPlan: ['', ],
      Status: ['', ]
        });
    }

    
    updateInspectionPlan(planNumber, revision, Item, Characteristics, SamplingPlan, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInspectionPlan(planNumber, revision, Item, Characteristics, SamplingPlan, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInspectionPlan']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInspectionPlan(params['id']).subscribe(res => {
                this.inspectionPlan = res;
            });
        });
    }
}