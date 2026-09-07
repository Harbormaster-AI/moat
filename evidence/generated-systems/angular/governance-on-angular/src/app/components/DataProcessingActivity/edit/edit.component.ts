import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataProcessingActivityService } from '../../../services/DataProcessingActivity.service';
import { SubBaseComponent } from '../../DataProcessingActivity/sub.base.component';


@Component({
    selector: 'app-edit-dataProcessingActivity',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataProcessingActivityComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataProcessingActivity';

    dataProcessingActivityForm: FormGroup;
    dataProcessingActivity: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataProcessingActivityService,
        private fb: FormBuilder
) {
        super(http);
        this.dataProcessingActivityForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      purpose: ['', Validators.required],
      startDate: ['', Validators.required],
      Organization: ['', ],
      DataCategories: ['', ],
      Systems: ['', ],
      Records: ['', ],
      PrivacyNotices: ['', ],
      ThirdParties: ['', ],
      Consents: ['', ],
      DataBreaches: ['', ],
      DataSubjectRequests: ['', ],
      LawfulBasis: ['', ]
        });
    }

    
    updateDataProcessingActivity(name, purpose, startDate, Organization, DataCategories, Systems, Records, PrivacyNotices, ThirdParties, Consents, DataBreaches, DataSubjectRequests, LawfulBasis): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataProcessingActivity(name, purpose, startDate, Organization, DataCategories, Systems, Records, PrivacyNotices, ThirdParties, Consents, DataBreaches, DataSubjectRequests, LawfulBasis, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataProcessingActivity']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataProcessingActivity(params['id']).subscribe(res => {
                this.dataProcessingActivity = res;
            });
        });
    }
}