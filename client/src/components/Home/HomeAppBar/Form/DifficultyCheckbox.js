import React, { useState } from "react";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import CheckBoxOutlineBlankIcon from "@material-ui/icons/CheckBoxOutlineBlank";
import CheckBoxIcon from "@material-ui/icons/CheckBox";
import Favorite from "@material-ui/icons/Favorite";
import FavoriteBorder from "@material-ui/icons/FavoriteBorder";

function DifficultyCheckbox({
  easyChecked,
  mediumChecked,
  advancedChecked,
  nativeChecked,
  handleChange,
}) {
  return (
    <FormGroup row>
      <FormControlLabel
        control={
          <Checkbox
            checked={easyChecked}
            onChange={handleChange}
            name="easyChecked"
            color="green"
          />
        }
        label="Easy"
      />
      <FormControlLabel
        control={
          <Checkbox
            checked={mediumChecked}
            onChange={handleChange}
            name="mediumChecked"
            color="primary"
          />
        }
        label="Medium"
      />
      <FormControlLabel
        control={
          <Checkbox
            checked={advancedChecked}
            onChange={handleChange}
            name="advancedChecked"
            color="secondary"
          />
        }
        label="Advanced"
      />
      <FormControlLabel
        control={
          <Checkbox
            checked={nativeChecked}
            onChange={handleChange}
            name="nativeChecked"
            color="default"
          />
        }
        label="Native"
      />
    </FormGroup>
  );
}

export default DifficultyCheckbox;
