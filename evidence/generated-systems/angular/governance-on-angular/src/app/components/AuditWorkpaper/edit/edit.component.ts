import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AuditWorkpaperService } from '../../../services/AuditWorkpaper.service';
import { SubBaseComponent } from '../../AuditWorkpaper/sub.base.component';


@Component({
    selector: 'app-edit-auditWorkpaper',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAuditWorkpaperComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AuditWorkpaper';

    auditWorkpaperForm: FormGroup;
    auditWorkpaper: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AuditWorkpaperService,
        private fb: FormBuilder
) {
        super(http);
        this.auditWorkpaperForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  workpaperRef: ['', Validators.required],
      subject: ['', Validators.required],
      workpaperUrl: ['', Validators.required],
      Engagement: ['', ],
      Evidence: ['', ],
      Findings: ['', ]
        });
    }

    
    updateAuditWorkpaper(workpaperRef, subject, workpaperUrl, Engagement, Evidence, Findings): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAuditWorkpaper(workpaperRef, subject, workpaperUrl, Engagement, Evidence, Findings, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAuditWorkpaper']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAuditWorkpaper(params['id']).subscribe(res => {
                this.auditWorkpaper = res;
            });
        });
    }
}