
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCertificationComponent } from './index.component';
import { CertificationService } from '../../../services/Certification.service';

describe('IndexCertificationComponent', () => {
  let component: IndexCertificationComponent;
  let fixture: ComponentFixture<IndexCertificationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCertificationComponent
      ],
      providers: [
        CertificationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCertificationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});