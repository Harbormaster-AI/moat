
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateGeoRegionComponent } from './create.component';
import { GeoRegionService } from '../../../services/GeoRegion.service';
import { Router } from '@angular/router';

describe('CreateGeoRegionComponent', () => {
  let component: CreateGeoRegionComponent;
  let fixture: ComponentFixture<CreateGeoRegionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateGeoRegionComponent
      ],
      providers: [
        GeoRegionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateGeoRegionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});