import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MRPRunService } from '../../../services/MRPRun.service';
import { SubBaseComponent } from '../../MRPRun/sub.base.component';


@Component({
    selector: 'app-edit-mRPRun',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMRPRunComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MRPRun';

    mRPRunForm: FormGroup;
    mRPRun: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MRPRunService,
        private fb: FormBuilder
) {
        super(http);
        this.mRPRunForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  runNumber: ['', Validators.required],
      runDateTime: ['', Validators.required],
      planningHorizonDays: ['', Validators.required],
      Plant: ['', ],
      PlannedOrders: ['', ],
      Status: ['', ]
        });
    }

    
    updateMRPRun(runNumber, runDateTime, planningHorizonDays, Plant, PlannedOrders, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMRPRun(runNumber, runDateTime, planningHorizonDays, Plant, PlannedOrders, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMRPRun']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMRPRun(params['id']).subscribe(res => {
                this.mRPRun = res;
            });
        });
    }
}