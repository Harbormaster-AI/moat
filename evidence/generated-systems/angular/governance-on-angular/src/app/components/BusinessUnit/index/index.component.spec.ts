
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBusinessUnitComponent } from './index.component';
import { BusinessUnitService } from '../../../services/BusinessUnit.service';

describe('IndexBusinessUnitComponent', () => {
  let component: IndexBusinessUnitComponent;
  let fixture: ComponentFixture<IndexBusinessUnitComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBusinessUnitComponent
      ],
      providers: [
        BusinessUnitService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBusinessUnitComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});