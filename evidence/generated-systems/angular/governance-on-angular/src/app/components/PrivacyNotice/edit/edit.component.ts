import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PrivacyNoticeService } from '../../../services/PrivacyNotice.service';
import { SubBaseComponent } from '../../PrivacyNotice/sub.base.component';


@Component({
    selector: 'app-edit-privacyNotice',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPrivacyNoticeComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PrivacyNotice';

    privacyNoticeForm: FormGroup;
    privacyNotice: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PrivacyNoticeService,
        private fb: FormBuilder
) {
        super(http);
        this.privacyNoticeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      audience: ['', Validators.required],
      versionLabel: ['', Validators.required],
      publicationDate: ['', Validators.required],
      publicationUrl: ['', Validators.required],
      ProcessingActivities: ['', ],
      Organization: ['', ],
      Consents: ['', ],
      Status: ['', ]
        });
    }

    
    updatePrivacyNotice(title, audience, versionLabel, publicationDate, publicationUrl, ProcessingActivities, Organization, Consents, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePrivacyNotice(title, audience, versionLabel, publicationDate, publicationUrl, ProcessingActivities, Organization, Consents, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPrivacyNotice']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPrivacyNotice(params['id']).subscribe(res => {
                this.privacyNotice = res;
            });
        });
    }
}