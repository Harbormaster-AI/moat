
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexKYCProfileComponent } from './index.component';
import { KYCProfileService } from '../../../services/KYCProfile.service';

describe('IndexKYCProfileComponent', () => {
  let component: IndexKYCProfileComponent;
  let fixture: ComponentFixture<IndexKYCProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexKYCProfileComponent
      ],
      providers: [
        KYCProfileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexKYCProfileComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});