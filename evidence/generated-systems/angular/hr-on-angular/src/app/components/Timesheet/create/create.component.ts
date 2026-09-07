import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TimesheetService } from '../../../services/Timesheet.service';
import { Timesheet } from '../../../models/Timesheet';
import { SubBaseComponent } from '../../Timesheet/sub.base.component';

@Component({
    selector: 'app-create-timesheet',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTimesheetComponent extends SubBaseComponent implements OnInit {

    title = 'Add Timesheet';

    timesheetForm: FormGroup;
    timesheet: Timesheet;

    constructor( http: HttpClient,
        private timesheetService: TimesheetService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTimesheet(periodStart, periodEnd, submissionDate, Employee, TimeEntries, Approvals, Status): void {
        this.timesheetService
        .addTimesheet(periodStart, periodEnd, submissionDate, Employee, TimeEntries, Approvals, Status)
            .subscribe(() => {
                this.router.navigate(['/indexTimesheet']);
            });
    }

    ngOnInit(): void {
    }
}