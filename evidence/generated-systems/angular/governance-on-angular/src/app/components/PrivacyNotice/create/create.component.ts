import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PrivacyNoticeService } from '../../../services/PrivacyNotice.service';
import { PrivacyNotice } from '../../../models/PrivacyNotice';
import { SubBaseComponent } from '../../PrivacyNotice/sub.base.component';

@Component({
    selector: 'app-create-privacyNotice',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePrivacyNoticeComponent extends SubBaseComponent implements OnInit {

    title = 'Add PrivacyNotice';

    privacyNoticeForm: FormGroup;
    privacyNotice: PrivacyNotice;

    constructor( http: HttpClient,
        private privacyNoticeService: PrivacyNoticeService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPrivacyNotice(title, audience, versionLabel, publicationDate, publicationUrl, ProcessingActivities, Organization, Consents, Status): void {
        this.privacyNoticeService
        .addPrivacyNotice(title, audience, versionLabel, publicationDate, publicationUrl, ProcessingActivities, Organization, Consents, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPrivacyNotice']);
            });
    }

    ngOnInit(): void {
    }
}