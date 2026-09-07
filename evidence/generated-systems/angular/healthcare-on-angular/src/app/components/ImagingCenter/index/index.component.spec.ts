
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexImagingCenterComponent } from './index.component';
import { ImagingCenterService } from '../../../services/ImagingCenter.service';

describe('IndexImagingCenterComponent', () => {
  let component: IndexImagingCenterComponent;
  let fixture: ComponentFixture<IndexImagingCenterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexImagingCenterComponent
      ],
      providers: [
        ImagingCenterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexImagingCenterComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});