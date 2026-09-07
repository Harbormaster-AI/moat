import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { Record_Service } from '../../../services/Record_.service';
import { SubBaseComponent } from '../../Record_/sub.base.component';


@Component({
    selector: 'app-edit-record_',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRecord_Component extends SubBaseComponent implements OnInit {

    title = 'Edit Record_';

    record_Form: FormGroup;
    record_: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: Record_Service,
        private fb: FormBuilder
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

    
    updateRecord_(title, creationDate, Repository, RetentionSchedule, ProcessingActivities, DataCategories, LegalHolds, DataSubjectRequests, RecordType, Classification, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRecord_(title, creationDate, Repository, RetentionSchedule, ProcessingActivities, DataCategories, LegalHolds, DataSubjectRequests, RecordType, Classification, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRecord_']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRecord_(params['id']).subscribe(res => {
                this.record_ = res;
            });
        });
    }
}