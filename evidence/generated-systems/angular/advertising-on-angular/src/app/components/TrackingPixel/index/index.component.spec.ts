
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTrackingPixelComponent } from './index.component';
import { TrackingPixelService } from '../../../services/TrackingPixel.service';

describe('IndexTrackingPixelComponent', () => {
  let component: IndexTrackingPixelComponent;
  let fixture: ComponentFixture<IndexTrackingPixelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTrackingPixelComponent
      ],
      providers: [
        TrackingPixelService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTrackingPixelComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});