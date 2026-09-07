import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MRPRunService } from '../../../services/MRPRun.service';
import { MRPRun } from '../../../models/MRPRun';
import { SubBaseComponent } from '../../MRPRun/sub.base.component';

@Component({
    selector: 'app-create-mRPRun',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMRPRunComponent extends SubBaseComponent implements OnInit {

    title = 'Add MRPRun';

    mRPRunForm: FormGroup;
    mRPRun: MRPRun;

    constructor( http: HttpClient,
        private mRPRunService: MRPRunService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMRPRun(runNumber, runDateTime, planningHorizonDays, Plant, PlannedOrders, Status): void {
        this.mRPRunService
        .addMRPRun(runNumber, runDateTime, planningHorizonDays, Plant, PlannedOrders, Status)
            .subscribe(() => {
                this.router.navigate(['/indexMRPRun']);
            });
    }

    ngOnInit(): void {
    }
}