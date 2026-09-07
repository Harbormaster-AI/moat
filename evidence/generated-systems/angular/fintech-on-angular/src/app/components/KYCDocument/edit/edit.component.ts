import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { KYCDocumentService } from '../../../services/KYCDocument.service';
import { SubBaseComponent } from '../../KYCDocument/sub.base.component';


@Component({
    selector: 'app-edit-kYCDocument',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditKYCDocumentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit KYCDocument';

    kYCDocumentForm: FormGroup;
    kYCDocument: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: KYCDocumentService,
        private fb: FormBuilder
) {
        super(http);
        this.kYCDocumentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reference: ['', Validators.required],
      issuedCountry: ['', Validators.required],
      expirationDate: ['', Validators.required],
      KycProfile: ['', ],
      DocumentType: ['', ],
      Status: ['', ]
        });
    }

    
    updateKYCDocument(reference, issuedCountry, expirationDate, KycProfile, DocumentType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateKYCDocument(reference, issuedCountry, expirationDate, KycProfile, DocumentType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexKYCDocument']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getKYCDocument(params['id']).subscribe(res => {
                this.kYCDocument = res;
            });
        });
    }
}