
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAerospaceManufacturerComponent } from './index.component';
import { AerospaceManufacturerService } from '../../../services/AerospaceManufacturer.service';

describe('IndexAerospaceManufacturerComponent', () => {
  let component: IndexAerospaceManufacturerComponent;
  let fixture: ComponentFixture<IndexAerospaceManufacturerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAerospaceManufacturerComponent
      ],
      providers: [
        AerospaceManufacturerService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAerospaceManufacturerComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});