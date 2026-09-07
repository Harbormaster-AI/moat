
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexExposureComponent } from './index.component';
import { ExposureService } from '../../../services/Exposure.service';

describe('IndexExposureComponent', () => {
  let component: IndexExposureComponent;
  let fixture: ComponentFixture<IndexExposureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexExposureComponent
      ],
      providers: [
        ExposureService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexExposureComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});