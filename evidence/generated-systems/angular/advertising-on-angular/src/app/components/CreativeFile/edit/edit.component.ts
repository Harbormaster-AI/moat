import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CreativeFileService } from '../../../services/CreativeFile.service';
import { SubBaseComponent } from '../../CreativeFile/sub.base.component';


@Component({
    selector: 'app-edit-creativeFile',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCreativeFileComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CreativeFile';

    creativeFileForm: FormGroup;
    creativeFile: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CreativeFileService,
        private fb: FormBuilder
) {
        super(http);
        this.creativeFileForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  uri: ['', Validators.required],
      fileSizeKB: ['', Validators.required],
      mimeType: ['', Validators.required],
      checksum: ['', Validators.required],
      CreativeAsset: ['', ]
        });
    }

    
    updateCreativeFile(uri, fileSizeKB, mimeType, checksum, CreativeAsset): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCreativeFile(uri, fileSizeKB, mimeType, checksum, CreativeAsset, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCreativeFile']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCreativeFile(params['id']).subscribe(res => {
                this.creativeFile = res;
            });
        });
    }
}