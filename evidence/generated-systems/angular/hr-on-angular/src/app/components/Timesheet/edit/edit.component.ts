import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TimesheetService } from '../../../services/Timesheet.service';
import { SubBaseComponent } from '../../Timesheet/sub.base.component';


@Component({
    selector: 'app-edit-timesheet',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTimesheetComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Timesheet';

    timesheetForm: FormGroup;
    timesheet: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TimesheetService,
        private fb: FormBuilder
) {
        super(http);
        this.timesheetForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  periodStart: ['', Validators.required],
      periodEnd: ['', Validators.required],
      submissionDate: ['', Validators.required],
      Employee: ['', ],
      TimeEntries: ['', ],
      Approvals: ['', ],
      Status: ['', ]
        });
    }

    
    updateTimesheet(periodStart, periodEnd, submissionDate, Employee, TimeEntries, Approvals, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTimesheet(periodStart, periodEnd, submissionDate, Employee, TimeEntries, Approvals, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTimesheet']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTimesheet(params['id']).subscribe(res => {
                this.timesheet = res;
            });
        });
    }
}