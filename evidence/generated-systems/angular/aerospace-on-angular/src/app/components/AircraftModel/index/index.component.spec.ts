
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAircraftModelComponent } from './index.component';
import { AircraftModelService } from '../../../services/AircraftModel.service';

describe('IndexAircraftModelComponent', () => {
  let component: IndexAircraftModelComponent;
  let fixture: ComponentFixture<IndexAircraftModelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAircraftModelComponent
      ],
      providers: [
        AircraftModelService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAircraftModelComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});