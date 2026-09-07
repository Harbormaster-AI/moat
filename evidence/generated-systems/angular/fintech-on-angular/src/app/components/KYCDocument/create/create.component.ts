import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { KYCDocumentService } from '../../../services/KYCDocument.service';
import { KYCDocument } from '../../../models/KYCDocument';
import { SubBaseComponent } from '../../KYCDocument/sub.base.component';

@Component({
    selector: 'app-create-kYCDocument',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateKYCDocumentComponent extends SubBaseComponent implements OnInit {

    title = 'Add KYCDocument';

    kYCDocumentForm: FormGroup;
    kYCDocument: KYCDocument;

    constructor( http: HttpClient,
        private kYCDocumentService: KYCDocumentService,
        private fb: FormBuilder,
        private router: Router
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

    
    addKYCDocument(reference, issuedCountry, expirationDate, KycProfile, DocumentType, Status): void {
        this.kYCDocumentService
        .addKYCDocument(reference, issuedCountry, expirationDate, KycProfile, DocumentType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexKYCDocument']);
            });
    }

    ngOnInit(): void {
    }
}