import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataSubjectRequestService } from '../../../services/DataSubjectRequest.service';
import { DataSubjectRequest } from '../../../models/DataSubjectRequest';
import { SubBaseComponent } from '../../DataSubjectRequest/sub.base.component';

@Component({
    selector: 'app-create-dataSubjectRequest',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataSubjectRequestComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataSubjectRequest';

    dataSubjectRequestForm: FormGroup;
    dataSubjectRequest: DataSubjectRequest;

    constructor( http: HttpClient,
        private dataSubjectRequestService: DataSubjectRequestService,
        private fb: FormBuilder,
        private router: Router
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

    
    addDataSubjectRequest(receivedDate, dueDate, requesterCountry, Organization, ProcessingActivities, Records, RequestType, Status): void {
        this.dataSubjectRequestService
        .addDataSubjectRequest(receivedDate, dueDate, requesterCountry, Organization, ProcessingActivities, Records, RequestType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexDataSubjectRequest']);
            });
    }

    ngOnInit(): void {
    }
}