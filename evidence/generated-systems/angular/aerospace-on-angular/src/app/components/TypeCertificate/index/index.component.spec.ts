
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTypeCertificateComponent } from './index.component';
import { TypeCertificateService } from '../../../services/TypeCertificate.service';

describe('IndexTypeCertificateComponent', () => {
  let component: IndexTypeCertificateComponent;
  let fixture: ComponentFixture<IndexTypeCertificateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTypeCertificateComponent
      ],
      providers: [
        TypeCertificateService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTypeCertificateComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});