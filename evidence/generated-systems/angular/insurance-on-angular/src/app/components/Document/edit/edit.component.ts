import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DocumentService } from '../../../services/Document.service';
import { SubBaseComponent } from '../../Document/sub.base.component';


@Component({
    selector: 'app-edit-document',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDocumentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Document';

    documentForm: FormGroup;
    document: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DocumentService,
        private fb: FormBuilder
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

    
    updateDocument(fileName, uploadedDate, Policy, Claim, Application, Customer, DocumentType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDocument(fileName, uploadedDate, Policy, Claim, Application, Customer, DocumentType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDocument']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDocument(params['id']).subscribe(res => {
                this.document = res;
            });
        });
    }
}