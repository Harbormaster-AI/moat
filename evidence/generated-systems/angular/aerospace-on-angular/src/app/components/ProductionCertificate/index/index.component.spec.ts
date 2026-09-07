
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexProductionCertificateComponent } from './index.component';
import { ProductionCertificateService } from '../../../services/ProductionCertificate.service';

describe('IndexProductionCertificateComponent', () => {
  let component: IndexProductionCertificateComponent;
  let fixture: ComponentFixture<IndexProductionCertificateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexProductionCertificateComponent
      ],
      providers: [
        ProductionCertificateService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexProductionCertificateComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});