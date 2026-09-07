import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TimeEntryService } from '../../../services/TimeEntry.service';
import { TimeEntry } from '../../../models/TimeEntry';
import { SubBaseComponent } from '../../TimeEntry/sub.base.component';

@Component({
    selector: 'app-create-timeEntry',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTimeEntryComponent extends SubBaseComponent implements OnInit {

    title = 'Add TimeEntry';

    timeEntryForm: FormGroup;
    timeEntry: TimeEntry;

    constructor( http: HttpClient,
        private timeEntryService: TimeEntryService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTimeEntry(entryDate, hoursWorked, Timesheet, Employee, CostCenter, EntryType): void {
        this.timeEntryService
        .addTimeEntry(entryDate, hoursWorked, Timesheet, Employee, CostCenter, EntryType)
            .subscribe(() => {
                this.router.navigate(['/indexTimeEntry']);
            });
    }

    ngOnInit(): void {
    }
}