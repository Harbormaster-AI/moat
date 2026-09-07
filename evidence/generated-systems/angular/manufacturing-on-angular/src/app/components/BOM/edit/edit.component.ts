import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BOMService } from '../../../services/BOM.service';
import { SubBaseComponent } from '../../BOM/sub.base.component';


@Component({
    selector: 'app-edit-bOM',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBOMComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BOM';

    bOMForm: FormGroup;
    bOM: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BOMService,
        private fb: FormBuilder
) {
        super(http);
        this.bOMForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  bomNumber: ['', Validators.required],
      revision: ['', Validators.required],
      effectivityStart: ['', Validators.required],
      effectivityEnd: ['', Validators.required],
      ParentItem: ['', ],
      BomItems: ['', ],
      Status: ['', ]
        });
    }

    
    updateBOM(bomNumber, revision, effectivityStart, effectivityEnd, ParentItem, BomItems, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBOM(bomNumber, revision, effectivityStart, effectivityEnd, ParentItem, BomItems, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBOM']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBOM(params['id']).subscribe(res => {
                this.bOM = res;
            });
        });
    }
}