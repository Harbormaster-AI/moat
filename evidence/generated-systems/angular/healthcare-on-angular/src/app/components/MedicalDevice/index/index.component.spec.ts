
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMedicalDeviceComponent } from './index.component';
import { MedicalDeviceService } from '../../../services/MedicalDevice.service';

describe('IndexMedicalDeviceComponent', () => {
  let component: IndexMedicalDeviceComponent;
  let fixture: ComponentFixture<IndexMedicalDeviceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMedicalDeviceComponent
      ],
      providers: [
        MedicalDeviceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMedicalDeviceComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});