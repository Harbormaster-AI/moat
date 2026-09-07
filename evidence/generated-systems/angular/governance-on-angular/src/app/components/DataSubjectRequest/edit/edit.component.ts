import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataSubjectRequestService } from '../../../services/DataSubjectRequest.service';
import { SubBaseComponent } from '../../DataSubjectRequest/sub.base.component';


@Component({
    selector: 'app-edit-dataSubjectRequest',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataSubjectRequestComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataSubjectRequest';

    dataSubjectRequestForm: FormGroup;
    dataSubjectRequest: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataSubjectRequestService,
        private fb: FormBuilder
) {
        super(http);
        this.dataSubjectRequestForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  receivedDate: ['', Validators.required],
      dueDate: ['', Validators.required],
      requesterCountry: ['', Validators.required],
      Organization: ['', ],
      ProcessingActivities: ['', ],
      Records: ['', ],
      RequestType: ['', ],
      Status: ['', ]
        });
    }

    
    updateDataSubjectRequest(receivedDate, dueDate, requesterCountry, Organization, ProcessingActivities, Records, RequestType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataSubjectRequest(receivedDate, dueDate, requesterCountry, Organization, ProcessingActivities, Records, RequestType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataSubjectRequest']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataSubjectRequest(params['id']).subscribe(res => {
                this.dataSubjectRequest = res;
            });
        });
    }
}