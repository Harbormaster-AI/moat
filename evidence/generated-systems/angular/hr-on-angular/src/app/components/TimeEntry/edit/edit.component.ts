import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TimeEntryService } from '../../../services/TimeEntry.service';
import { SubBaseComponent } from '../../TimeEntry/sub.base.component';


@Component({
    selector: 'app-edit-timeEntry',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTimeEntryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TimeEntry';

    timeEntryForm: FormGroup;
    timeEntry: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TimeEntryService,
        private fb: FormBuilder
) {
        super(http);
        this.timeEntryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  entryDate: ['', Validators.required],
      hoursWorked: ['', Validators.required],
      Timesheet: ['', ],
      Employee: ['', ],
      CostCenter: ['', ],
      EntryType: ['', ]
        });
    }

    
    updateTimeEntry(entryDate, hoursWorked, Timesheet, Employee, CostCenter, EntryType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTimeEntry(entryDate, hoursWorked, Timesheet, Employee, CostCenter, EntryType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTimeEntry']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTimeEntry(params['id']).subscribe(res => {
                this.timeEntry = res;
            });
        });
    }
}