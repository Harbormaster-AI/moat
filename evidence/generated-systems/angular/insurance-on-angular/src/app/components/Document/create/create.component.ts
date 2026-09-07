import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DocumentService } from '../../../services/Document.service';
import { Document } from '../../../models/Document';
import { SubBaseComponent } from '../../Document/sub.base.component';

@Component({
    selector: 'app-create-document',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDocumentComponent extends SubBaseComponent implements OnInit {

    title = 'Add Document';

    documentForm: FormGroup;
    document: Document;

    constructor( http: HttpClient,
        private documentService: DocumentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.documentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  fileName: ['', Validators.required],
      uploadedDate: ['', Validators.required],
      Policy: ['', ],
      Claim: ['', ],
      Application: ['', ],
      Customer: ['', ],
      DocumentType: ['', ]
        });
    }

    
    addDocument(fileName, uploadedDate, Policy, Claim, Application, Customer, DocumentType): void {
        this.documentService
        .addDocument(fileName, uploadedDate, Policy, Claim, Application, Customer, DocumentType)
            .subscribe(() => {
                this.router.navigate(['/indexDocument']);
            });
    }

    ngOnInit(): void {
    }
}