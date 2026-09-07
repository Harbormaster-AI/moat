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
                  name: ['', Validators.required],
      fileUrl: ['', Validators.required],
      uploadedDate: ['', Validators.required],
      Candidate: ['', ],
      Employee: ['', ],
      DocumentType: ['', ]
        });
    }

    
    updateDocument(name, fileUrl, uploadedDate, Candidate, Employee, DocumentType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDocument(name, fileUrl, uploadedDate, Candidate, Employee, DocumentType, params['id'])
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