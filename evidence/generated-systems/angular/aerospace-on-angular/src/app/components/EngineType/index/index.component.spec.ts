
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEngineTypeComponent } from './index.component';
import { EngineTypeService } from '../../../services/EngineType.service';

describe('IndexEngineTypeComponent', () => {
  let component: IndexEngineTypeComponent;
  let fixture: ComponentFixture<IndexEngineTypeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEngineTypeComponent
      ],
      providers: [
        EngineTypeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEngineTypeComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});