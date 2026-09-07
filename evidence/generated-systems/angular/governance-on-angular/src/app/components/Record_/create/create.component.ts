import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Record_Service } from '../../../services/Record_.service';
import { Record_ } from '../../../models/Record_';
import { SubBaseComponent } from '../../Record_/sub.base.component';

@Component({
    selector: 'app-create-record_',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRecord_Component extends SubBaseComponent implements OnInit {

    title = 'Add Record_';

    record_Form: FormGroup;
    record_: Record_;

    constructor( http: HttpClient,
        private record_Service: Record_Service,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.record_Form = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      creationDate: ['', Validators.required],
      Repository: ['', ],
      RetentionSchedule: ['', ],
      ProcessingActivities: ['', ],
      DataCategories: ['', ],
      LegalHolds: ['', ],
      DataSubjectRequests: ['', ],
      RecordType: ['', ],
      Classification: ['', ],
      Status: ['', ]
        });
    }

    
    addRecord_(title, creationDate, Repository, RetentionSchedule, ProcessingActivities, DataCategories, LegalHolds, DataSubjectRequests, RecordType, Classification, Status): void {
        this.record_Service
        .addRecord_(title, creationDate, Repository, RetentionSchedule, ProcessingActivities, DataCategories, LegalHolds, DataSubjectRequests, RecordType, Classification, Status)
            .subscribe(() => {
                this.router.navigate(['/indexRecord_']);
            });
    }

    ngOnInit(): void {
    }
}