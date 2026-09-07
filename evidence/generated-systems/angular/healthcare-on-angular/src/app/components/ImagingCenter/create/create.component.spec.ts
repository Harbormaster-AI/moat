
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateImagingCenterComponent } from './create.component';
import { ImagingCenterService } from '../../../services/ImagingCenter.service';
import { Router } from '@angular/router';

describe('CreateImagingCenterComponent', () => {
  let component: CreateImagingCenterComponent;
  let fixture: ComponentFixture<CreateImagingCenterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateImagingCenterComponent
      ],
      providers: [
        ImagingCenterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateImagingCenterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});