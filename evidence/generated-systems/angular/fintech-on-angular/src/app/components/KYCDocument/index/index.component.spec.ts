
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexKYCDocumentComponent } from './index.component';
import { KYCDocumentService } from '../../../services/KYCDocument.service';

describe('IndexKYCDocumentComponent', () => {
  let component: IndexKYCDocumentComponent;
  let fixture: ComponentFixture<IndexKYCDocumentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexKYCDocumentComponent
      ],
      providers: [
        KYCDocumentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexKYCDocumentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});